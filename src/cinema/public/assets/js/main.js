$( document ).ready(function() {
    
    var selectedSeat = [];
    var span = document.getElementsByClassName("close")[0];
    var span2 = document.getElementsByClassName("close")[1];

    $( '.seats' ).click(function() {
        if($(this).hasClass('empty-seat')){
            selectedSeat.push($(this).attr('data-id'));
            
            $(this).removeClass('empty-seat').addClass('selected-seat');
        } else if($(this).hasClass('booked-seat')){
            if(selectedSeat.length == 0){

                getBooking($(this).attr('data-id'));
                $('#cancelBookingModal').show();
            }
        } else if($(this).hasClass('selected-seat')){
            var id = $(this).attr('data-id');
            selectedSeat.splice(selectedSeat.indexOf(id), 1);
            console.log(selectedSeat);
            $(this).removeClass('selected-seat').addClass('empty-seat');
        }
        //show booking button if seat is selected
        if(selectedSeat.length > 0) {
            $('#bookTkt').show();
        } else {
            $('#bookTkt').hide();
        }
    });

    $( '#bookBtn' ).click(function() {
        saveBooking();
    });

    $( '.closeModal' ).click(function() {
        location.reload();
    });

    //save booking information
    function saveBooking(){
        var name = $('#name').val();
        var email = $('#email').val();
        $.ajax({
            method: "POST",
            url: "/save-booking",
            dataType: 'JSON',
            data: {id : JSON.stringify(selectedSeat), name : name, email : email}
        })
        .done(function( res ) {
            if(res.success == "ok"){
                $('#response').removeClass('err-res').addClass('success-res');
                $('.close').hide();
                $('.close-modal').show();
                $('#bookingForm').hide();
            }else{
                $('#response').removeClass('success-res').addClass('err-res');
            }
            
            
            $('#response').html(res.message);
            $('#response').show();
        });
    }

    //get booking information
    function getBooking(seatId){
        $('#seatNos').html('');
        $.ajax({
            method: "POST",
            url: "/get-booking",
            dataType: 'JSON',
            data: {id : seatId}
        })
        .done(function( res ) {
            if(res.success == "ok"){
                for (i=0; i<res.seats.length;i++) { 
                    var lbl = $('<label>').text('# '+res.seats[i]);
                    var chk = $('<input/>').attr({ type: 'checkbox', name:'cseat', value: res.seats[i]}).addClass("chk");
                    $("#seatNos").append(lbl);
                    $("#seatNos").append(chk);

                } 
            }else{
                
            }
        });
    }

    $( '#cancelBookBtn' ).click(function() {

        var cancelSeat = [];
        $.each($("input[name='cseat']:checked"), function(){
            cancelSeat.push($(this).val());
        });
        var email = $('#cemail').val();
        $.ajax({
            method: "POST",
            url: "/cancel-booking",
            dataType: 'JSON',
            data: {ids : JSON.stringify(cancelSeat), email : email}
        })
        .done(function(res) {
            if(res.success == "ok"){
                $('#cancelRes').addClass('success-res');
            }else{
                $('#cancelRes').addClass('err-res');
            }

            $('#cclose').hide();
            $('#closeCancelModal').show();
            $('#cform').hide();
            $('#cancelRes').html(res.message);
            $('#cancelRes').show();
        });
    });

    $( '#bookTkt' ).click(function() {
        $('#name, #email, #cemail').val('');
        $('#myModal').show();
    });

    span.onclick = function() {
      $('#myModal').hide();
    }

    span2.onclick = function() {
      $('#cancelBookingModal').hide();
    }
 });