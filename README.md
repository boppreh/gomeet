gomeet
======


Go server to punch NAT holes. The application that wants to listen makes a
GET request to `http://gomeet/meet/some_unique_id`, notes the port the request
went through (this has to be done locally) and starts listening on that port.
The client then requests the same page and gets the server port and address,
so it can start a connection on that address.

The trick is that the port the server is listening on is different from the
port the world sees, because of NAT. By communicating to the server it is
possible to know which private port corresponds to which public port, so
the connecting client can establish a connection even with both sides
are behind NAT.
