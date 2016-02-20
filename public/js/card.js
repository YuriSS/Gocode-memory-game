game['card'] = function() {

    'use strict'

    var Model = {
        send_move: function(index) {
            var data = {id: index}
            $.ajax({
                url: "api/move",
                method: "POST",
                data: data,
                dataType: "json",
                success: function(res) {
                    console.log(res);
                }
            });
        }
    };

    var Controller = {
        init: function() {
            View.init();
            this.bind_events();
        },

        bind_events: function() {
            View.stage.on('click', function(event) {
                var el = $(event.target).closest('li');
                var index = $('li').index(el);
                Model.send_move(index);
            });
        }
    };

    var View = {
        init: function() {
            this.stage = $('#stage');
        }
    };

    return {
        init: function() {
            Controller.init();
            return this;
        }
    };

};
