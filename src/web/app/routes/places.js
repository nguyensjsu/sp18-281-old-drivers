var express = require('express');
var router = express.Router();
var Place = require('../models/place');
var ObjectId = require('mongodb').ObjectId;

router.get('/about', function(req, res, next) {


        res.render('about',{});

});

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

router.get('/test', function(req, res, next) {

    Place.findOne({'name': 'Balboa Park'}, function(err, doc){
        res.render('search', {items: doc, user:req.user});
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
