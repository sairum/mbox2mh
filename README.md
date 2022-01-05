# mbox2mh


## Convert thunderbird MBOX structure into MH structure to be read by ClawsMail

*mbox2mh* is a small go utility to convert thunderbird MBOX (Mailbox) folder structure for Local mail storage of emails into a MH (Message Handling System) folder structure capable of being read by other email clients such as *claws-mail*. It should not be used directly in the *INBOX* folder of thunderbird!

The current version is capable of reading MBOX folders (usually directories terminating in a  ".sbd" extension , and recursively create all necessary directories and splitting the MBOX files into as many individual email files needed, named numerically, starting from one ("1") to N.

