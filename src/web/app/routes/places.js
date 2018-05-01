var express = require('express');
var router = express.Router();
// var Place = require('../models/place');
// var ObjectId = require('mongodb').ObjectId;
var Request = require("request");


router.get('/about', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/3b0fa0dc-6e35-4cb4-bd6f-cfaf71ef9c13", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('about',{inventory: result});
    });

});


router.get('/product', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('product',{americano: result});
    });

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('product',{latte: result});
    });

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('product',{mocha: result});
    });

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('product',{cappuccino: result});
    });

});


router.get('/americano', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{americano: result});
    });

});

router.get('/latte', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{latte: result});
    });

});

router.get('/mocha', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{mocha: result});
    });

});

router.get('/cappuccino', function(req, res, next) {

    Request.get("http://localhost:8080/inventory/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('detail',{cappuccino: result});
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
