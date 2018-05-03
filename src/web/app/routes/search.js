var express = require('express');
var router = express.Router();
var Place = require('../models/place');


router.post('/search', function(req, res) {
    var keyWord = req.body.key;
    console.log(JSON.stringify(req.body));
    var regex = new RegExp(["^", keyWord, "$"].join(""), "i");

    Place.find({'keywords': regex}, function(err, doc){
        console.log(regex);
        console.log(doc);
        console.log('Got through');
        res.render('search', {results: doc, keyWord: keyWord});
    });
});


/* router.post('/insert', function(req, res, next) {
    var item = {
        name: req.body.name,
        content: req.body.content
    };

    var data = new UserReview(item);
    data.save();

    res.redirect('/display');
});

router.post('/update', function(req, res, next) {
    var id = req.body.id;

    UserReview.findById(id, function(err, doc) {
        if (err) {
            console.error('error, no entry found');
        }
        doc.name = req.body.name;
        doc.content = req.body.content;
        doc.save();
    });
    res.redirect('/display');
});

router.post('/delete', function(req, res, next) {
    //console.log(JSON.stringify(req));
    var id = req.body.id;
    UserReview.findByIdAndRemove(id).exec();
    res.redirect('/display');
}); */

module.exports = router;
