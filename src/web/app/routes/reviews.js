var express = require('express');
var router = express.Router();
var UserReview = require('../models/review');


router.get('/display', function(req, res, next) {

    UserReview.find({}, function(err, doc){
        res.render('display', {items: doc, user:req.user});
    });
});

router.post('/insert', function(req, res, next) {
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
});

module.exports = router;
