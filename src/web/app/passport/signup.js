var LocalStrategy   = require('passport-local').Strategy;
// var User = require('../models/user');
var bCrypt = require('bcrypt-nodejs');
var Request = require('request');

module.exports = function(passport){

	passport.use('signup', new LocalStrategy({
            passReqToCallback : true // allows us to pass back the entire request to the callback
        },
        function(req, name, password, done) {

            name = req.body.name
            phone = req.body.phone
            balance = req.body.balance

            Request.post('http://localhost:8080/user?name={name}&phone={phone}&balance={balance}', function(err, response, body) {
                if (!error && response.statusCode == 200) {
                    console.log(body) // Print the google web page.
                }
                console.dir(JSON.parse(body))

                res.render('index', {})
            })


            /*
            Request.post({
                "headers": { "content-type": "application/json" },
                "url": "http://user/",
                "body": JSON.stringify({
                    "Name": req.body.name,
                    "Phone": req.body.phone
                    "Balance": req.body.balance
                })}, (error, response, body) => {
                if(error) {
                    return console.dir(error);
                }
                console.dir(JSON.parse(body));


                });

            res.redirect('/login')
            */


            /*
            findOrCreateUser = function(){
                // find a user in Mongo with provided username
                User.findOne({ 'username' :  username }, function(err, user) {
                    // In case of any error, return using the done method
                    if (err){
                        console.log('Error in SignUp: '+err);
                        return done(err);
                    }
                    // already exists
                    if (user) {
                        console.log('User already exists with username: '+username);
                        return done(null, false, req.flash('message','User Already Exists'));
                    } else {
                        // if there is no user with that email
                        // create the user
                        var newUser = new User();

                        // set the user's local credentials
                        newUser.username = username;
                        newUser.password = createHash(password);
                        newUser.email = req.param('email');
                        newUser.firstName = req.param('firstName');
                        newUser.lastName = req.param('lastName');

                        // save the user
                        newUser.save(function(err) {
                            if (err){
                                console.log('Error in Saving user: '+err);  
                                throw err;  
                            }
                            console.log('User Registration succesful');    
                            return done(null, newUser);
                        });
                    }
                });
            };*/
            // Delay the execution of findOrCreateUser and execute the method
            // in the next tick of the event loop
            // process.nextTick(findOrCreateUser);
        })
    );

    /* Generates hash using bCrypt
    var createHash = function(password){
        return bCrypt.hashSync(password, bCrypt.genSaltSync(10), null);
    }*/

}