# Stairs Backend Challenge

This is my solution for the Stairs Backend Challenge.

Written in Go, using MongoDB and Go Modules mode. 


## Install and Run

This API uses [Go](http://golang.org) and [MongoDB](https://mongodb.com).

The MongoDB is instanced using [Docker Compose](https://github.com/docker/compose). 

### Setting up the Database
The MongoDB is instanced using docker-compose. 

Navigate to the API folder and raise the database with:
```sh
$ docker-compose-up
```

### Creating the database and a collection

Once MongoDB is up, you must create a database and a collection. 
This setup also raises a **Mongo Express** instance to make this easier.

The address to the Mongo Express configuration page is: **`localhost:8081`**

### Run the API 
After raising the database instance, run the API:
```sh
$ go run main.go
```

## API Description

This is a REST API made with Go that exposes 5 endpoints:

<hr>
<div style="background-color: #faf7f1">
<div style="width: auto; height: 35px; background: #dffed8">
<code style="display: inline-block; width: 70px; position: relative; top: 3px;font-size: 15px;font-weight: bolder; background-color: lightgreen; color: black; margin-left: 7px; padding: 3px 17px;">
POST
</code>
<code style="font-size: 15px;background:inherit; position: relative; top: 4px;">
/wines
</code>
<span style="font-size: 13px;background:inherit; position: relative; top: 4px;">
Get all wines
</span>
</div>
</div>
<br>

*Parameters:* 
```
No parameters
```

<hr>

<div style="width: auto; height: 35px; background: #e2f3fc">
<code style="width: 70px; display: inline-block; position: relative; top: 3px;font-size: 15px;font-weight: bolder; background-color: #78cffe; color: black; margin-left: 7px; padding: 3px 20px;">
GET
</code>
<code style="font-size: 15px;background:inherit; position: relative; top: 4px;">
/wines
</code>
<span style="font-size: 13px;background:inherit; position: relative; top: 4px;">
Get all wines
</span>
</div>
<br>

*Parameters:* 

```
No parameters
```

<hr>

<div style="width: auto; height: 35px; background: #dffed8">
<code style="display: inline-block; width: 70px; position: relative; top: 3px;font-size: 15px;font-weight: bolder; background-color: lightgreen; color: black; margin-left: 7px; padding: 3px 17px;">
POST</code>
<code style="font-size: 15px;background:inherit; position: relative; top: 4px;">
/wines/new
</code>
<span style="font-size: 13px;background:inherit; position: relative; top: 4px;">
Add wines
</span>
</div>
<br>

*Parameters:* 

<table style="font-size: 13px;border: none" cellspacing="0" cellpadding="0">
<tr>
<td>
name</td><td> string</td>
</tr>
<tr>
<td>harvest</td><td>integer</td>
</tr>
<tr>
<td>country</td><td>string</td>
</tr>
<tr>
<td>amount</td><td>int</td>
</tr>
<tr>
<td>description</td><td>string</td> 
</tr>
</table>


*Usage Example:*

```json
{
  "name": "Petrus",
  "harvest": 1984,
  "country": "France",
  "ammount": 0,
  "description": "This wine comes from the regions of..."
}
```

<div style="clear:both">
NOTE: MongoDB adds the ObjectID automatically if its not specified
</div>

<hr>

<div style="clear: both;width: auto; height: 35px; background: #fae3b7">
<code style="width: 70px; display: inline-block; position: relative; top: 3.5px;font-size: 15px;font-weight: bolder; background-color: #f2ab2e; color: black; margin-left: 7px; padding: 3px 21px;">
PUT</code>
<code style="font-size: 15px;background:inherit; position: relative; top: 4px;">
/wines/{id}
</code>
<span style="font-size: 13px;background:inherit; position: relative; top: 4px;">
Update a wine
</span>
</div>
<br>

*Parameters:* 

<table style="font-size: 13px;border: none" cellspacing="0" cellpadding="0">
<tr>
<td>id</td><td>int</td>
</tr>
<tr>
<td>wine</td><td>Struct</td>
</tr>
</table>


*Usage Example:*

```json
{
  "id": "5d881cf1d663f81498131a81",
  "wine": {
    "name": "New Name",
    "harvest": 2005,
    "country": "New Country",
    "ammount": 0,
    "description": "New description for the wine..."
  }
}
```

<hr>
<div style="width: auto; height: 35px; background: #fce2e2">
<code style="width: 80px; display: inline-block; position: relative; top: 3.5px;font-size: 15px;font-weight: bolder; background-color: #fa5353; color: black; margin-left: 7px; padding: 3px 12px;">
DELETE</code>
<code style="font-size: 15px;background:inherit; position: relative; top: 4px;">
/wines/{id}
</code>
<span style="font-size: 13px;background:inherit; position: relative; top: 4px;">
Deletes a wine
</span>
</div>
<br>

*Parameters:* 

<table style="font-size: 13px;border: none" cellspacing="0" cellpadding="0">
<tr>
<td>id</td><td>int</td>
</tr>
</table>
<br>

<hr>


