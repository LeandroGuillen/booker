#!/usr/bin/env node
 
// Examples:
// - Get Orion Context Broker version:
//		./booker.js
// - Get table1 booking status:
//		./booker.js --table table1
// - Set table1 booking status:
//		./booker.js --table table1 --booked true
 
var request = require('request');
var argv = require('optimist').argv;
 
var urlBase = "http://leandroguillen.com:1026"
 
if (argv.table && argv.booked)
{
	payload={
		    "contextElements": [
		        {
		          "type": "Table",
		          "isPattern": "false",
		          "id": argv.table,
		          "attributes": [
		            {
		              "name": "booked",
		              "type": "boolean",
		              "value": argv.booked
		            }
		          ]
		        }
		    ],
		    "updateAction": "UPDATE"
		}
	
	request.post(
	    urlBase+'/v1/updateContext',
	    { json: payload },
	    function (error, response, body) {
	        if (!error && response.statusCode == 200) {
	        	if (body.statusCode)
	        	{
	        		console.log(body.statusCode.reasonPhrase)
	        	}
	        	else
	        	{
		        	if (body.contextResponses[0].statusCode.code == "200")
		        	{
		        		console.log('Table "'+argv.table+'": booked state is set to "'+argv.booked+'"')
		            	//console.log(JSON.stringify(body, null, 2))
		        	}
		        	else
		        	{
		        		console.log(body.contextResponses[0].statusCode.reasonPhrase)
		        	}		
	        	}
	        } else {
	        	console.log(error)
	        }
	    }
	);
}
else if (argv.table)
{
	var options = {
	    url: urlBase+'/v1/contextEntities/' + argv.table + '/attributes/booked',
	    headers: {
	        'Accept': 'application/json'
	    }
	};
	request.get(
	    options,
	    function (error, response, body) {
	        if (!error && response.statusCode == 200) {
	        	var json = JSON.parse(body)
	            if (json.statusCode && !json.attributes)
	        	{
	        		console.log(json.statusCode)
	        	}
	        	else
	        	{
		            var isBooked = json.attributes[0].value
		            if (isBooked === 'true')
		            {
		            	console.log('Table "'+argv.table+'" is booked.')
		            }
		            else
		            {
		            	console.log('Table "'+argv.table+'" is free.')
		            }
	            }
	        } else {
	        	console.log('error:' + error)
	        	console.log('statusCode:' + response.statusCode)
	        }
	    }
	);
}
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
