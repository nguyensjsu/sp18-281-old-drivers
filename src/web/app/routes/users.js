var express = require('express');
var router = express.Router();

var isAuthenticated = function (req, res, next) {
    // if user is authenticated in the session, call the next() to call the next request handler
    // Passport adds this method to request object. A middleware is allowed to add properties to
    // request and response objects
    if (req.isAuthenticated())
        return next();
    // if the user is not authenticated then redirect him to the login page
    return next();
};


module.exports = function(passport){

    router.get('/', isAuthenticated,
        function(req, res)    //routing for homepage
        {
            res.render('homepage', { user: req.user });
        });

	/* GET login page. */
	router.get('/login', isAuthenticated ,function(req, res) {
    	// Display the Login page with any flash message, if any
        if(req.isAuthenticated())
            res.redirect('/home');
        else
            res.render('index', { message: req.flash('message') });
	});

	/* GET Registration Page */
	router.get('/signup', isAuthenticated, function(req, res){
	    if(req.isAuthenticated())
	        res.redirect('/home');
        else
            res.render('register',{message: req.flash('message')});
	});

	/* Handle Registration POST */
	router.post('/signup', passport.authenticate('signup', {
		successRedirect: '/home',
		failureRedirect: '/signup',
		failureFlash : true  
	}));

    /* Handle Login POST */
    router.post('/login', passport.authenticate('login', {
        successRedirect: '/',
        failureRedirect: '/login',
        failureFlash : true
    }));

	/* GET Home Page */
	router.get('/home', isAuthenticated, function(req, res){
	    if(req.isAuthenticated())
	        res.render('home', {user: req.user})
        else
            res.redirect('/')
    });



	/* Handle Logout */
	router.get('/signout', function(req, res) {
		req.logout();
		res.redirect('/');
	});



	return router;
};




