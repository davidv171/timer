# Timer

## Why ?

Testing go-pls

## Why is this on github

Go doesn't let me just have stuff locally...

## I wanna use this

No you don't. Script it:
Countdown:

```bash

#!/bin/bash

for (( c=$1; c>=0; c-- ))
do
	echo -e \\033c &&\
		echo -n "$c." | figlet -c -k &&\
		sleep 1;
done


```


Timer:

```bash

#!/bin/bash

for (( c=0; c<=$1; c++ ))
do
	echo -e \\033c &&\
		echo -n "$c." | figlet -c -k &&\
		sleep 1;
done


```
