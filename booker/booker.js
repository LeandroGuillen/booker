#!/usr/bin/env node
 
// Examples:
// - Get Orion Context Broker version:
//		./booker.js
// - Create a table named 'table1':
//		./booked.js --table table1 --create
// - Get table1 booking status:
//		./booker.js --table table1
// - Set table1 booking status:
//		./booker.js --table table1 --booked true
// - Delete a table named 'table1':
//		./booked.js --table table1 --delete
 
var request = require('request');
var argv = require('optimist').argv;
 
var urlBase = "http://localhost:1026"
 
// Changing booking status
if (argv.table && (argv.booked || argv.book || argv.unbook))
{
	if (argv.book) {
		newStatus = "true"
	} else if (argv.unbook) {
		newStatus = "false"
	} else {
		newStatus = argv.booked
	}

	payload={
		   "booked": newStatus
	}
		 
	request.patch(
	    urlBase+'/v2/entities/' + argv.table,
	    { json: payload },
	    function (error, response, body) {
	        if (!error && response.statusCode == 204) {
       			console.log('Table "'+argv.table+'": booked state is set to "'+newStatus+'"')
	        } else {
	        	console.log("Maybe the table does not exist?")
	        }
	    }
	);
}

// Creating a new table
else if (argv.table && argv.create)
{
	payload={
	    "id": argv.table,
	    "type": "Table",
	    "isPattern": "false",
	    "booked": "false"
	}
	
	request.post(
	    urlBase+'/v2/entities',
	    { json: payload },
	    function (error, response, body) {
	        if (!error && response.statusCode == 201) {
	        	console.log('Table "'+argv.table+'" was created.')
	        } else {
	        	console.log(error)
	        }
	    }
	);
}

// Removing a table
else if (argv.table && argv.del)
{
	var options = {
	    url: urlBase+'/v2/entities/' + argv.table,
	    headers: {
	        'Accept': 'application/json'
	    }
	};

	request.del(
	    options,
	    function (error, response, body) {
            if (!error && response.statusCode == 204) {
	        	console.log('Table "'+argv.table+'" was deleted.')
	        } else {
	        	console.log("Maybe the table does not exist already?")
	        }
	    }
	);
}

// Querying booking status
else if (argv.table)
{
	var options = {
	    url: urlBase+'/v2/entities/' + argv.table,
	    headers: {
	        'Accept': 'application/json'
	    }
	};
	request.get(
	    options,
	    function (error, response, body) {
	        if (!error && response.statusCode == 200)
	        {
	        	var json = JSON.parse(body)
	            var isBooked = json.booked
	            if (isBooked === "true")
	            {
	            	console.log('Table "'+argv.table+'" is booked.')
	            }
	            else
	            {
	            	console.log('Table "'+argv.table+'" is free.')
	            }
	        } else if (!error && response.statusCode == 404)
	        {
	        	console.log("The table does not exist.")
	        }
	        else {
	        	console.log('error:' + error)
	        	console.log('statusCode:' + response.statusCode)
	        	console.log("Maybe the table does not exist?")
	        }
	    }
	);
}

// By default just print Orion version
else
{
	console.log('No parameters found')
	
	var options = {
	    url: urlBase+'/version',
	    headers: {
	        'Accept': 'application/json'
	    }
	};
	
	request.get(
	    options,
	    function (error, response, body) {
	        if (!error && response.statusCode == 200) {
	            console.log(body)
	        }
	    }
	);
}
