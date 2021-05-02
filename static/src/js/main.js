import $ from 'jquery';

$('#card').keyup( function(e) {
    let endpoint = 'http://localhost:8090/api'
    this.value = this.value.replace(/[^0-9\.]/g,'');
    this.value = this.value.replace(/-/g,'');
    var value = $("#card").val();
    $.ajax({
        url: endpoint + "?card=" + value,
        contentType: "application/json",
        dataType: 'json',
        success: function(result){
            if(result.Valid && value.length > 11){
                $("#luhn").html("<i class=\"material-icons\">check</i>");
                $("#mii").html(result.MII);
                $("#iin").html(result.Issuer);
                $("#pan").html(result.PAN);
            } else {
                $("#luhn").html("<i class=\"material-icons red-text\">close</i>");
                $("#mii").html("<i class=\"material-icons red-text\">close</i>");
                $("#iin").html("<i class=\"material-icons red-text\">close</i>");
                $("#pan").html("<i class=\"material-icons red-text\">close</i>");
            }
            console.log(result);
        }
    })


});
