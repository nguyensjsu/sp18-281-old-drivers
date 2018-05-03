var express = require('express');
var router = express.Router();
//var Place = require('../models/place');


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


// latte        a6764ac4-6fc8-4826-8247-58e25909c80d
// americano    5d85f3eb-8576-44e3-b075-bc28129cbd8f
// mocha        3dadd5bb-456d-43c0-9b49-f50cfb091826
// cappuccino   f3ef10b1-b937-458b-9ee2-5a90b1ca0936    

module.exports = router;
