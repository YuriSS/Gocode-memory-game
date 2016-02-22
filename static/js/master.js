game['card'] = function() {

    'use strict'

    var Model = {
        send_move: function(index, callback) {
            var data = JSON.stringify({index: index});
            $.ajax({
                url: "api/move/",
                method: "POST",
                data: data,
            })
                .done(callback);
        }
    };

    var Controller = {
        init: function() {
            View.init();
            this.bind_events();
        },

        bind_events: function() {
            var _this = this;
            View.stage.on('click', function(event) {
                var el = $(event.target).closest('li');
                var index = $('li').index(el);
                Model.send_move(index, _this.new_play(el));
            });
        },

        new_play: function(el) {
            return function(data) {
                View.turn_image(el, data.image);
            }
        }
    };

    var View = {
        init: function() {
            this.stage = $('#stage');
        },

        turn_image: function(el, image) {
            el.find('img').attr('src', image);
        }
    };

    return {
        init: function() {
            Controller.init();
            return this;
        }
    };

};
game['start'] = (function() {

    'use strict'

    var card = game.card().init();

})();
