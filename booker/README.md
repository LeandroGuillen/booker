# booker
Simple Orion Context Broker Node app

## Dependencies
* request
* optimist

## Examples:
Get Orion Context Broker version:

```
node booker.js
```

Get table1 booking status:

```
node booker.js --table table1
```

Set table1 to booked:

```
node booker.js --table table1 --booked true
```

OR

```
node booker.js --table table1 --book
```

Set table1 to unbooked:

```
node booker.js --table table1 --booked false
```

OR

```
node booker.js --table table1 --unbook
```

Create a new table:

```
node booker.js --table table1 --create
```

Delerte a table:

```
node booker.js --table table1 --del
```