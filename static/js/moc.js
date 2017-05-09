class MOC {
    formatTime(t) {
        var time = moment(t)
        return time.format("YYYY-MM-DD HH:mm:ss")
    }
    toggleLoading(flag){
        if (flag){
            console.log("loading")
            $("#loading").css("display","block")
            console.log($("#loading").css("display"))
        }else{
            console.log("loading finished")
            $("#loading").css("display","none")
            console.log($("#loading").css("display"))
        }
    }
}

var moc = new(MOC);

