package main

import (
    "fmt"
    "flag"
    "strings"
    "os"
    "path/filepath"
    "io/ioutil"
    "log"
    "bufio"
//    "time"
    "strconv"
)

var outfile *os.File

/*
 * Simple function to spit out how to use the program!
 * Two directories, a source and a destination, are mandatory
 * and should not be the same!
 */

func usage() {
    fmt.Fprintf(os.Stderr, "usage: mbox2mh [-test] <src directory> <dest directory>\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func processFile( fileName, destdir string, test bool) {

    /*
     * Check if this is a Thunderbird index file (".msf")
     * If it is, skip it. Otherwise, it should be a .mbox
     * file. We will create a directory with the same name
     * to drop individual mail files there!
     */

    if filepath.Ext(fileName) != ".msf" {

        path := filepath.Join(destdir,filepath.Base(fileName))
        fmt.Printf("Write %s into %s\n", fileName, path )

        /*
         * We create a directory with the name of
         * the current mbox (
         */

        err := os.MkdirAll(path, 0755)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Problems writing folder '%s': %s\n", path, err.Error())
            //return err
        }

        // // To use this code enable import package time
        //start := time.Now()

        /*
         * We can now open the .mbox file which has the same
         * name as the new destination directory
         */

        file, err := os.Open(fileName)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        /*
         * Start scanning the mbox file line by line
         */

        scanner := bufio.NewScanner(file)

        /*
         * individual mails will be stored in text files
         * numbered from 1 to the maximum number of mesages
         * (mails)
         */

        var i = 0;

        for scanner.Scan() {
            line := scanner.Text()

            /*
             * Check if this line starts with "From - ...". If so, it's a
             * begining of an email header. Check if 'i' is greater than 0.
             * If it is, it means that a file had been already open for
             * writing so close it and increment 'i'. The variable 'i' holds
             * the name of a file to write mail text to.
             *
             * Note that the '-' on the header may be different. Apparently,
             * some mail clients write 'MAILER-DEAMON' instead. AFAIK, current
             * versions of thunderbird don't use this signature but if you get
             * strange results it may be because of this. Worst than this,
             * apparently some non-english clients use the worf 'DE' instead
             * of 'From'! Go figure...
             */

            if strings.HasPrefix(line, "From - ") {
                // This is a lazy check. In .mbox format, mails start with a
                // line of the form
                //
                // "From - Tue Jan 10 22:54:05 2017"
                //
                // so the first line should have seven (7) tokens
                // with the last one denoting a year which should be
                // greater than 1970 the year of the beginning of
                // UNIX time.
                s := strings.Split(line, " ")

                // The year is the 7th slice (index starts in 0, so it's 6)
                year, _ := strconv.Atoi(s[6])

                if ( ( year > 1969 ) && (s[1] == "-" ) ) {
                    if ( i > 0 ) && !test {
                        outfile.Close()
                    }
                    i++
                    f := filepath.Join(path,strconv.Itoa(i))

                    if !test {
                        outfile, err = os.Create(f)
                    } else {
                        fmt.Printf("Writing %s\n", f)
                    }
                    if err != nil {
                        //return err
                    }
                }

            }
            //fmt.Printf(line)
            if !test {
                _, _ = outfile.WriteString(line+"\n")
            }
        }

        outfile.Close()
        // // To use this code enable import package time
        // duration := time.Since(start)
        // fmt.Printf("duration was %f seconds\n",duration.Seconds())

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
}


/*
 * Process each entry in the directory hierarchy
 */

func process_dir(srcdir, destdir string, test bool) error {

    items, err := ioutil.ReadDir(srcdir)
    if err != nil {
        // Probably not a directory! Let's try to process
        // it as a file
        processFile( srcdir, destdir, test )
    } else {

        for _, item := range items {

            src := filepath.Join(srcdir,item.Name())

            /*
            * Check if the current item is a directory. If it is
            * it should have an extension ".sbd" which should be
            * removed
            */

            if item.IsDir() {
                dirname := strings.TrimSuffix(item.Name(), ".sbd")
                newdestdir := filepath.Join(destdir,dirname)
                process_dir(src, newdestdir, test)
            } else {
                processFile( src, destdir, test )
            }
        }
    }
    return nil
}

func main() {

    /*
     * Check if we are testing only!
     */

    testPtr := flag.Bool("t", false, "a bool, if used no files will be written")

    flag.Usage = usage
    flag.Parse()

    args := flag.Args()

    if len(args) < 1 {
        usage()
        //fmt.Println("Source and destination directories are missing.")
        os.Exit(1)
    }

    if len(args) < 2 {
        fmt.Println("Destination directory is missing.")
        os.Exit(1)
    }

    /*
     * Save source and destination directories
     */

    srcdir  := args[0]
    destdir := args[1]

    /*
     * Check if source and destination directories are the same
     */

    if args[0] == args[1] {
        fmt.Println("Source and destination directories are similar and should be different!")
        os.Exit(1)
    }

    /*
     * Check if source directory exists!
     */

    if _, err := os.Stat(srcdir); os.IsNotExist(err) {
        fmt.Fprintf(os.Stderr,"Dource directory '%s' does not exist!\n", srcdir)
        os.Exit(1)
    }


    _ = process_dir(srcdir, destdir, *testPtr)

    return

}
