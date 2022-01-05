# mbox2mh


### Convert thunderbird MBOX structure into MH structure to be read by ClawsMail

*mbox2mh* is a small go utility to convert thunderbird MBOX (Mailbox) folder structure for Local mail storage of emails into a MH (Message Handling System) folder structure capable of being read by other email clients such as *claws-mail*. It should not be used directly in the *INBOX* folder of thunderbird!

The current version is capable of reading MBOX folders (usually directories terminating in a  ".sbd" extension), and recursively create all necessary directories and splitting the MBOX files into as many individual email files as needed, named numerically, starting from one ("1") to N.

Consider the following structure

```bash
~/Mail
  ├── Inbox
  ├── Inbox.msf
  ├── Junk.msf
  ├── Sent
  ├── Sent.msf
  ├── Templates.msf
  ├── Trash
  ├── Trash.msf
  ├── friends.sbd
  │   ├── maria
  │   ├── maria.msf
  │   ├── joana.sbd
  │   │   ├── pictures
  │   │   └── pictures.msf
  │   ├── ted.sbd
  │   │   ├── photos
  │   │   ├── photos.msg
  │   │   └── photos.sbd
  │   │       ├── summer
  │   │       └── summer.msf
  │   ├── other
  │   └── other.msf
  ├── home
  └── home.msf
```

The command

```bash
$ mbox2mh ~/Mail/friends /home/username/Friends
```

will produce the following structure in the folder Friends located at the home directory of *username*

```bash
~/Friends
  └── friends
      ├── maria
      │   ├── 1
      │   ├── 2
      │   ├── 3
      │   ...
      ├── joana
      │   ├── 1
      │   ├── 2
      │   ├── ...
      │   └── pictures
      │       ├── 1
      │       ├── 2
      │       ├── 3
      │       ...
      ├── ted
      │   ├── 1
      │   ├── 2
      │   ├── 3
      │   ├── ...
      │   └── photos
      │       ├── 1
      │       ├── 2
      │       ├── 3
      │       ├── ...
      │       └── summer
      │           ├── 1
      │           ├── 2
      │           ├── 3
      │           ...
      └── other
          ├── 1
          ├── 2
          ├── 3
          ...

```
