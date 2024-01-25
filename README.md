# Deployhelper
An mini internal tool I developed for fun because I wanted to learn Go and I also wanted to learn how to automate my tasks more.
There is no better way to learn a language but to build something useful with it.
 Could I have done this with a bash script? Yes, but I wanted to learn Go and really learn by shooting myself in the foot :).
## What does it do?
It is a command line tool that basically checks out your Dev branch on the repo you are working on, pulls changes from your 
branch to Dev and push it up. That's it, basic and nothing special.

## How to use it?
First, make sure you have Go installed on your machine. If you don't, you can download it [here](https://golang.org/dl/).

Also, make sure your ssh-agent is running and you have added your ssh key to it.

Clone this repository and run `go build` to build the binary.
Then. you can place the binary in a folder and set the PATH variable to that folder.

