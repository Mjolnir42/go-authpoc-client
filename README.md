Client/Server-Authentication POC Example in go - Client
=======================================================

This is the client part of my POC playtest authentication mockup.
To play around with some classes and test ideas.

It currently mocks the following actions:

1. register user account
2. authenticate for token
3. validate an existing token

Register user account
---------------------

This will prompt for username and password to be used and try to
register that combination with the server. On sucess, a Token as well
as the token expiry time will be printed out.

```
go-authpoc-client register

~% ./go-authpoc-client register
Enter username: foo
Enter password:
Repeat password:
✔ Entered passwords match
Password score    (0-4): 4
Estimated entropy (bit): 61.175000
Estimated time to crack: centuries
Select this password? (y/n): n
Enter password:
Repeat password:
✔ Entered passwords match
Password score    (0-4): 4
Estimated entropy (bit): 61.175000
Estimated time to crack: centuries
Select this password? (y/n): y
Token: 68f9698fe2540c525fe35b15c6ae1a1788e079962b2ada3d1872c7665c95e148
Expires at: 2016-04-23T21:40:22Z
```

Authenticate for Token
----------------------

```
go-authpoc-client authenticate [-u|--user user]

~% ./go-authpoc-client authenticate
Enter username: foo
Enter password:
Token: 1ec5c269d7becd7f194b527951f8403ca1ef32926537dcc390dc27b1003badee
Expires at: 2016-04-24T00:27:11Z
```

This will prompt for username, if not set via cli flag, and password
and then try to request a token with those credentials.
On success, a Token as well as the token expiry time will be printed
out.

Validate token
--------------

```
go-authpoc-client validate [-u|--user user] [-t|--token token]

~% ./go-authpoc-client validate
Enter username: foo
Enter token:
✔ Token successfully verified

~% ./go-authpoc-client validate
Enter username: foo
Enter token:
✘ Verification failed: Response code was: 401 Unauthorized
```

This will prompt for username and token if they have not been supplied
via cli flags. Then a protected ressource is requested from the
server.

Remarks
=======

At any time, the master branches of client and server may or may not
compile or be able to talk to each other. Client and server built from
the same tag should however.

TLS Setup
=========

The server certificate must be copied into the client's working
directory as `./server.pem`.
