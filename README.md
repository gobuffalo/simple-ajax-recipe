# Simple Ajax Recipe

This simple [Buffalo](https://gobuffalo.io) application shows how you can use the out of the box features of the [Buffalo asset pipeline](https://gobuffalo.io/en/docs/assets) to quickly add Ajax to an application.

## Installation and Setup

* Make sure you have the latest version of [Buffalo](https://gobuffalo.io) install and working.
* Make sure your [Node](https://nodejs.org/en/) and [Yarn](https://yarnpkg.com/en/) installations are relatively up to date.

```bash
$ go get -u -v github.com/gobuffalo/simple-ajax-recipe
$ cd $GOPATH/github.com/gobuffalo/simple-ajax-recipe
$ buffalo setup
```

## Running

Run the following command to start the Buffalo dev server and then navigate to [http://localhost:3000](http://localhost:3000). It might take a second for Node to compile the front end stuff. Sorry about that.

```bash
$ buffalo dev
```

## jQuery UJS

The heart of all this is in the [https://github.com/rails/jquery-ujs/wiki](https://github.com/rails/jquery-ujs/wiki) plug-in that ships with Buffalo by default.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
