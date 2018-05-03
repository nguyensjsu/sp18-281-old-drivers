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

    var min  = date.getMinutes();
    min = (min < 10 ? "0" : "") + min;

    var sec  = date.getSeconds();
    sec = (sec < 10 ? "0" : "") + sec;

    var year = date.getFullYear();

    var month = date.getMonth() + 1;
    month = (month < 10 ? "0" : "") + month;

    var day  = date.getDate();
    day = (day < 10 ? "0" : "") + day;


    var time = year + ":" + month + ":" + day + ":" + hour + ":" + min + ":" + sec;
    res.render('order', {time: time})

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
        res.render('places_test', {items: doc, user:req.user});
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


/*
router.get('/seattle', function(req, res, next) {

    Place.findOne({'name': 'Seattle'}, function(err, doc){
        res.render('cities_test', {items: doc, user:req.user});
    });
});

router.get('/san-diego', function(req, res, next) {

    Place.findOne({'name': 'San Diego'}, function(err, doc){
        res.render('cities_test', {items: doc, user:req.user});
    });
});

router.get('/san-francisco', function(req, res, next) {

    Place.findOne({'name': 'San Francisco'}, function(err, doc){
        res.render('cities_test', {items: doc, user:req.user});
    });
});

router.get('/los-angeles', function(req, res, next) {

    Place.findOne({'name': 'Los Angeles'}, function(err, doc){
        res.render('cities_test', {items: doc, user:req.user});
    });
});

router.get('/golden-gate-bridge', function(req, res, next) {

    Place.findOne({'name': 'Golden Gate Bridge'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/alcatraz', function(req, res, next) {

    Place.findOne({'name': 'Alcatraz'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/lake-tahoe', function(req, res, next) {

    Place.findOne({'name': 'Lake Tahoe'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/yosemite-national-park', function(req, res, next) {

    Place.findOne({'name': 'Yosemite National Park'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/napa-valley', function(req, res, next) {

    Place.findOne({'name': 'Napa Valley'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/mendocino-coast', function(req, res, next) {

    Place.findOne({'name': 'Mendocino Coast'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/uss-midway-museum', function(req, res, next) {

    Place.findOne({'name': 'USS Midway Museum'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});


router.get('/la-jolla-cove', function(req, res, next) {

    Place.findOne({'name': 'La Jolla Cove'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/universal-studios-hollywood', function(req, res, next) {

    Place.findOne({'name': 'Universal Studios Hollywood'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/disneyland-park', function(req, res, next) {

    Place.findOne({'name': 'Disneyland Park'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/griffith-observatory', function(req, res, next) {

    Place.findOne({'name': 'Griffith Observatory'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/six-flags-magic-mountain', function(req, res, next) {

    Place.findOne({'name': 'Six Flags Magic Mountain'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/chihuly-garden-and-glass', function(req, res, next) {

    Place.findOne({'name': 'Chihuly Garden and Glass'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/pike-place-market', function(req, res, next) {

    Place.findOne({'name': 'Pike Place Market'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/space-needle', function(req, res, next) {

    Place.findOne({'name': 'Space Needle'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/washington-state-ferries', function(req, res, next) {

    Place.findOne({'name': 'Washington State Ferries'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});

router.get('/san-diego-zoo', function(req, res, next) {

    Place.findOne({'name': 'San Diego Zoo'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});


router.get('/balboa-park', function(req, res, next) {

    Place.findOne({'name': 'Balboa Park'}, function(err, doc){
        res.render('places_test', {items: doc, user:req.user});
    });
});
*/

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
