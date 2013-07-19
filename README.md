goaws
=====

When Amazon doesn't support Go, we write our own classes ;)

LICENSE
-------
BSD

documentation
-------------
[package documentation at godoc.org](http://godoc.org/github.com/mthie/goaws)

or

[package documentation with code view at gowalker.org](http://gowalker.org/github.com/mthie/goaws)

build and install
=================

install from source
-------------------

Install [Go 1][3], either [from source][4] or [with a prepackaged binary][5].

Then run

go get github.com/mthie/goaws

see the examples for usage.

[3]: http://golang.org
[4]: http://golang.org/doc/install/source
[5]: http://golang.org/doc/install

documentation
-------------
        AccessKey := "YoUrAcCeSsKeY"
        SecretKey := "YoUrSeCrEtKeY"

        r := goaws.NewRoute53(AccessKey, SecretKey)
        zones := r.GetHostedZones()

contributing
============

Contributions are welcome. Please open an issue or send me a pull request for a dedicated branch.
Make sure the git commit hooks show it works.

git commit hooks
-----------------------
enable commit hooks via

        cd .git ; rm -rf hooks; ln -s ../git-hooks hooks ; cd ..

