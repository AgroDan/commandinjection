# Command Injection

This is the code I wrote for the command injection talk I made at SCCC on 4/16/25. This counts as two parts:

1. The Docker container, with has the Go code and template use for the webpage, and..
2. `attack.py`, which is the python code I wrote in class to showcase how to exploit this application.

The run this yourself, make sure that Docker is properly installed on your machine! This is easier to explain on a Linux box, so it is probably better to start from there.

First, build the container in this directory:

```terminal
$ sudo docker build -t cmi .
```

This should take a minute or two because it has to download the appropriate docker images, then compile the Go code. ***Note***: This may not run on a Macbook if it is using an ARM architecture. If you are one of the lucky few to have an x86-architecture Mac, then you should have no problem building this. But if it fails, that's why!

Now run the container like so:

```terminal
$ sudo docker run --rm -p 8085:8080 -d cmi
```

It should display a SHA256 hash if it was successful. Note that I am forwarding `localhost:8085` to the container's port `8080`. You can attempt to find the IP address of the running container yourself if you'd like!

Now you can execute `attack.py`. This should interact with the running container according to the steps listed above!

Good luck and happy hacking!

-Agr0
