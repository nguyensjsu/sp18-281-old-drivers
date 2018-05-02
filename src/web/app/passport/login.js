var LocalStrategy   = require('passport-local').Strategy;
// var User = require('../models/user');
var bCrypt = require('bcrypt-nodejs');

module.exports = function(passport){

	passport.use('login', new LocalStrategy({
            passReqToCallback : true },
            function(req, name, password, done) { 
                // call get user api, bind response with home page
                if (password = "1") {
                    Request.get("http://localhost:8080/user/{userid}", (error, response, body) => {
                        if(error) {
                           return console.dir(error);
                        }
                        console.dir(JSON.parse(body));

                        result = JSON.parse(body)

                        res.render('home',{user: result});


                    });
                }
            })
    );

    /*
    var isValidPassword = function(user, password){
        return bCrypt.compareSync(password, user.password);
    }
    */
}