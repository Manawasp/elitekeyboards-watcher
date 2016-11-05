# elitekeyboards-watcher

This project watch [elitekeybaords website](http://elitekeyboards.com) and notify by email when there is change inside of the keyboards stock (only TKL Topre).

### Why this project ?

I was tired to check keyboard availability every day on the EK website (I wanted to buy one of their keyboard), so I made this project for fun (one year before). It was my first go project and pretty happy about the result, after one year I'm still receiving notification. I cleaned the code a little.      
(Finaly I bought FC660C :) and maybe will buy a Topre TKL or FC980C soon).

## Installation

You need install dependencies libraries (used by gokogiri) :

```
# apt-get install dpk-config libxml2-dev
```
Then install go library by hand (go get) or use [glide](http://glide.sh)

```
$ glide install
```
You are ready to compile
```
$ make
```

### Licence

MIT Licence
