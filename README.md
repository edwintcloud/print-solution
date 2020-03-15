# print-solution
A side project to try making a small SaaS print service. This is a very rudementary example but basically we have three key components: desktop client, server, and a web dashboard. The desktop client will communicate with the server via TCP socket for this example. When a regular API request is sent from the web dashboard to the server (such as print some text on this available printer), the server will translate this to a TCP message with a print OPCODE sent to the corresponding desktop client. The desktop client will receive this request, pipe the input into an `lp` command to print to an installed local printer, and send the server back an ACK which is then relayed back to the web dashboard.

__Desktop Client__
- TCP server
- Send POST request to server on startup registering its address, port, and printer information

__Server__
- API
- POST register a new desktop client
- GET list of printers (from desktop client info registered in POST api) and current status (ACK to registered clients)
- POST create a new job for an available printer

__Web Dashboard__
- Link to install a desktop client (will register itself on startup)
- List of printers in data table with action to print (data table calls GET list api once per second to update status)
- On print action modal with text box appears, user can type in text and click print

### Technology Used
- desktop-client: Golang (stdlib packages)
- server: Golang (stdlib, Echo, MongoDB)
- web-dashboard: Quasar (SPA), vue.js