Autorization()

function Autorization(){
  if (localStorage.getItem('ton-connect-storage_bridge-connection') !== null) {
    window.location.href = 'https://176-99-11-185.cloudvps.regruhosting.ru/home';
  } else {
    window.location.href = 'https://176-99-11-185.cloudvps.regruhosting.ru/';
  }  
}

