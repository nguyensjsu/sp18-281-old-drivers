var mongoose = require('mongoose');

var UserReview = mongoose.model('UserReview',{
    name: {
        type: String,
        required: true
    },
    content: String,
    create_date: {
        type: Date,
        default: Date.now
    },
    update_date: {
        type: Date,
        default: Date.now
    }
});

module.exports = UserReview;