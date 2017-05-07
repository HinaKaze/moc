class MOC {
    getFormatTime(t) {
        var d = new Date();
        d.setTime(Date.parse(t));
        return d.getFullYear() + "/" + d.getMonth() + "/" + d.getDate() + " " + d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
    }
}

var moc = new(MOC);

Vue.component("timer",{
    props:["timestr"],
    template: '<span>{{ timestr }}</span>'
})

Vue.component("rtimer",{
    props:["timestr"],
    template: '<span>{{ timestr }}</span>'
})