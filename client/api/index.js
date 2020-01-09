const express = require('express');
const app = express();

// const PROTO_PATH = __dirname + '/../../server/sumpb/sum.proto';
// const SERVER_HOST = process.env.SERVER_HOST || 'localhost';
// const client = caller(`${SERVER_HOST}:60000`, PROTO_PATH, 'SumService');

var PROTO_PATH = __dirname + '/../../server/sumpb/sum.proto';
var grpc = require('grpc');
var sum_proto = grpc.load(PROTO_PATH).sumpb;
var client = new sum_proto.SumService('localhost:60000', grpc.credentials.createInsecure());

app.get('/rpc', (req, res) => {
    // console.log(Number(req.query.num1))
    // console.log(Number(req.query.num2))
    client.Sum({num1: Number(req.query.num1), num2: Number(req.query.num2)}, (err, response) => {
        res.send(response);
    });
});

module.exports = {
  path: '/api',
  handler: app
}
