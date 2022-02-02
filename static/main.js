const icmfi = new Vue({
  el: '#icmfi',
  data: {
    links: []
  },
  created () {
    fetch('/icmfi/links')
    .then(response => response.json())
    .then(json => {
      this.links = json.links
    })
  }
})

// https://api.myjson.com/bins/74l63
