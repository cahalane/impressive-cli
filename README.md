# DEPRECATED
Will not work with versions of Impressive following [v1.0.1](https://github.com/colm2/impressive/releases/tag/1.0.1).

UCC broke compatibility, so this program is essentially useless in its current form.

## impressive-cli

![impressive logo](http://i.imgur.com/dXS1iob.png)

A simple wrapper for [https://github.com/colm2/impressive].

To create an iCal file for your college timetable:

* Make sure your timetable is set up on [mytimetable.ucc.ie](https://mytimetable.ucc.ie)
* Have Go installed and on your PATH.
* `go install github.com/colm2/impressive-cli`
* `impressive-cli email@umail.ucc.ie password > file.ical`
