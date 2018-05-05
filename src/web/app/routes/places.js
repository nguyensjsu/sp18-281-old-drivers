var express = require('express');
var router = express.Router();
var Place = require('../models/place');
var ObjectId = require('mongodb').ObjectId;

router.get('/about', function(req, res, next) {

    res.render('about',{});

});

router.get('/order', function(req, res, next) {

    var date = new Date();

    var hour = date.getHours();
    hour = (hour < 10 ? "0" : "") + hour;
    /*
    var min  = date.getMinutes();
    min = (min < 10 ? "0" : "") + min;
    
    var sec  = date.getSeconds();
    sec = (sec < 10 ? "0" : "") + sec;
    */
    var year = date.getFullYear();

    var month = date.getMonth() + 1;
    month = (month < 10 ? "0" : "") + month;

    var day  = date.getDate();
    day = (day < 10 ? "0" : "") + day;


    var time = year + ":" + month + ":" + day + ":" + hour;// + ":" + min + ":" + sec;
    res.render('order', {time: time})

});

router.get('/post_order', function(req, res, next) {

    var date = new Date();

    var hour = date.getHours();
    hour = (hour < 10 ? "0" : "") + hour;
    /*
    var min  = date.getMinutes();
    min = (min < 10 ? "0" : "") + min;
    
    var sec  = date.getSeconds();
    sec = (sec < 10 ? "0" : "") + sec;
    */
    var year = date.getFullYear();

    var month = date.getMonth() + 1;
    month = (month < 10 ? "0" : "") + month;

    var day  = date.getDate();
    day = (day < 10 ? "0" : "") + day;


    var time = year + ":" + month + ":" + day + ":" + hour;// + ":" + min + ":" + sec;
    res.render('post_order', {time: time})

});

router.get('/product', function(req, res, next) {

    Place.findOne({'name': 'Product'}, function(err, doc){
        res.render('cities_test', {items: doc, user:req.user});
    });
});

router.get('/americano', function(req, res, next) {

    Place.findOne({'name': 'Americano'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/cappuccino', function(req, res, next) {

    Place.findOne({'name': 'Cappuccino'}, function(err, doc){
        res.render('places', {items: doc, user:req.user});
    });
});

router.get('/mocha', function(req, res, next) {

    Place.findOne({'name': 'Mocha'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/latte', function(req, res, next) {

    Place.findOne({'name': 'Latte'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});



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

        res.render('places_test', {items: updated, user:req.user});

    });

});

module.exports = router;
