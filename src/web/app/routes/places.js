var express = require('express');
var router = express.Router();
// var Place = require('../models/place');
// var ObjectId = require('mongodb').ObjectId;
var Request = require("request");


router.get('/about', function(req, res, next) {

        res.render('about',{});

});


router.get('/product', function(req, res, next) {

    var latte, americano, mocha, cappuccino

    // latte
    Request.get("http://18.144.40.71:8000/inventory/inventory/96d8492a-802f-4f0a-9b04-f82ef5a16e2f", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        latte = JSON.parse(body)

        
    });

    // americano
    Request.get("http://18.144.40.71:8000/inventory/inventory/5d85f3eb-8576-44e3-b075-bc28129cbd8f", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        americano = JSON.parse(body)

    });

    // mocha
    Request.get("http://18.144.40.71:8000/inventory/inventory/5feecca7-bbd5-494e-be2a-1ca4bd469d4e", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        mocha = JSON.parse(body)

    });

    // cappuccino
    Request.get("http://18.144.40.71:8000/inventory/inventory/f3ef10b1-b937-458b-9ee2-5a90b1ca0936 ", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        cappuccino = JSON.parse(body)

    });

    res.render('product', {latte: latte, americano: americano, mocha: mocha, cappuccino: cappuccino})

});


router.get('/americano', function(req, res, next) {

    /*
    Request.get("http://18.144.40.71:8000/inventory/inventory/5d85f3eb-8576-44e3-b075-bc28129cbd8f", function (error, response, body) {
        if(body == "Inventory not exist") {
           return console.dir("error");
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)*/

        res.render('detail',{});
      
    //});

});




router.get('/latte', function(req, res, next) {

    Request.get("http://18.144.40.71:8000/inventory/inventory/96d8492a-802f-4f0a-9b04-f82ef5a16e2f", function (error, response, body) {
        if(error) {
           return console.dir(error);
        }
        // console.dir(JSON.parse(body));

        // result = JSON.parse(body)

        res.render('detail',{inventory: body});
    });

});

router.get('/mocha', function(req, res, next) {

    Request.get("http://18.144.40.71:8000/inventory/inventory/5feecca7-bbd5-494e-be2a-1ca4bd469d4e", function (error, response, body) {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{inventory: result});
    });

});

router.get('/cappuccino', function(req, res, next) {

    Request.get("http://18.144.40.71:8000/inventory/inventory/f3ef10b1-b937-458b-9ee2-5a90b1ca0936", function (error, response, body) {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{inventory: result});
    });

});


// invoke add review api
router.post('/addReview', function (req, res, next) {
    console.log(req.body);

    Place.findOne({ "name": req.body.place }, function(err, doc) {

        var updated = doc;
        var newReview = {
            content: req.body.content,
            userName: req.user.username,
        };
        updated.reviews.push(newReview);
        updated.save();

        res.render('detail', {items: updated, user:req.user});

    });

});

module.exports = router;
