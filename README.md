# Net-Cat
This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.


## Usage 
Server:
```
go run . $port
```
Client:
```
nc $IP $port
```
## Some warnings
The length of nickname should be more than 2 and less than 20 


## Our team

 - [mkabyken](https://01.alem.school/git/mkabyken)