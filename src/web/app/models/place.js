var mongoose = require('mongoose');

var Place = mongoose.model('Place', {
    name: String,
    url: String,
    keywords: [String],
    info: [{
        type: String
    }],
    website: [String],
    description: String,
    images: [{
        type: String
    }],
    scenes: [{
        name: String,
        url: String
    }],
    reviews: [
        {
            userName: String,
            content: String,
            createDate: {type: Date, default: Date.now()},
            updateDate: {type: Date, default: Date.now()}
        }
    ]
});

module.exports = Place;