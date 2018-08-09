# dwbrite.com
A personal web server intended for hosting dwbrite.com
Hosts a home page and basic blog, as well as a portfolio.
It is designed to run without super user permissions, 
as a supplement to an existing web server configuration, such as nginx or apache.

## Roll your own
For the purposes of creating simple and concrete examples, we'll be using nginx.

* Clone the repository
* Set up a PostgreSQL database with the 3* tables as described in `examples/tables.sql`.
Make sure the service is running.
  * It is recommended to store the database password in the user's `~/.pgpass` file, or in environment variables.
* Add an nginx configuration routing to the port this program runs on.
  * See `examples/nginx.conf` if you intend on enabling SSL/TLS access.
  * Make sure the ports used by nginx are open and being forwarded to the server machine.
* Create and deploy your SSL/TLS certificates. I use Let's Encrypt + `certbot` for my certs.
  * You'll want to create deploy scripts which copy the required files with proper permissions 
  to the user running the go server.
* Modify the constants in `main.go` to suit your needs. That is, 
  PostgreSQL access data, certificate and key locations, and the port which the application will run.
* Install the `pq` PostgreSQL driver with `go get github.com/lib/pq`
* Finally, run the server!

## License
Copyright (C) 2018 Devin Brite  
SPDX-License-Identifier: GPL-3.0-or-later OR  CC0-1.0 OR MIT

Full texts can be found in `LICENSE.md`