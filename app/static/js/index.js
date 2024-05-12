Autorization()

function Autorization(){
  if (localStorage.getItem('ton-connect-storage_bridge-connection') === null) {
    window.location.replace('/')
  }
}

