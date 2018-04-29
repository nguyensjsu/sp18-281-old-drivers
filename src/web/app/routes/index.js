var express = require('express');
var router = express.Router();


router.get('/maps', function(req, res, next) {


        res.render('maps', {});

});

router.get('/bucketlist', function(req, res, next) {


    res.render('bucketlist', {});

});

router.get('/contact', function(req,res){
   res.render('contact', {user: req.user});
});

router.get('/places', function(req, res, next) {


    res.render('places', {});

});

router.get('/SD_intro', function(req, res, next) {


    res.render('../place/SD_intro', {});

});

router.get('/SD_food', function(req, res, next) {


    res.render('../place/SD_food', {});

});

router.get('/SD_tour', function(req, res, next) {


    res.render('../place/SD_tour', {});

});

module.exports = router;
