const icmfi = new Vue({
  el: '#icmfi',
  data: {
    links: []
  },
  created () {
    fetch('http://127.0.0.1:8080/icmfi/links')
    .then(response => response.json())
    .then(json => {
      this.links = json.links
    })
  }
})

// https://api.myjson.com/bins/74l63
const x5732 = new Vue({
  el: '#x5732',
  data: {
    MessageType: ""
  },
  created () {
    fetch('http://127.0.0.1:8080/0x5732')
    .then(response => response.json())
    .then(json => {
      this.MessageType = json.MessageType
    })
  }
})
