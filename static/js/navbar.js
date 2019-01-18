
function getAllUserInfo() {
    $.get("/getUserInfo",function(data,status){
        $('#div-content').html(data);
    });
}


function GetPageMessage(tp,page) {
    $.get("http://192.168.1.225:8000/"+tp+"?page="+page +"&pageSize=10",function(data,status){
        $('#user-table').html(data);
    });

}

function getOrderInfo() {
    
}

