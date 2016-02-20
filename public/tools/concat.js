var concat = require('concat');
var path = require('path');
var dir = path.join(__dirname, '../../static/js/');


var js_files = [
    '/js/card.js',
    '/js/main.js'
]


for (var i in js_files) {
    js_files[i] = path.join(__dirname, '..' + js_files[i])
}


concat(js_files, (dir + 'master.js'), function(err) {
    if(err) {
        return console.error('JS concatenation fail: ' + err);
    }
    return console.log("Concatenation js wins!!!");
});
