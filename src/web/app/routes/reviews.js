var express = require('express');
var router = express.Router();
var UserReview = require('../models/review');


router.get('/display', function(req, res, next) {

  Request.get("http://localhost:8080/review/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('places_test',{review: result});
    });


})


router.post('/insert', function(req, res, next) {

    Request.post({
        "headers": { "content-type": "application/json" },
        "url": "http://review/",
        "body": JSON.stringify({
            "content": req.body.content,
            "date": date.Now()
        })
    }, (error, response, body) => {
        if(error) {
            return console.dir(error);
        }
        console.dir(JSON.parse(body));


    });

    res.redirect('/display')


})

router.get('/update', function(req, res, next) {

  Request.get("http://localhost:8080/review/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('places_test',{review: result});
    });


})

router.get('/delete', function(req, res, next) {

  Request.get("http://localhost:8080/review/", (error, response, body) => {
        if(error) {
           return console.dir(error);
        }
        console.dir(JSON.parse(body));

        result = JSON.parse(body)

        res.render('places_test',{review: result});
    });


})




module.exports = router;
