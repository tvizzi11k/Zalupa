Autorization()
function Autorization(){
  if (localStorage.getItem('ton-connect-storage_bridge-connection') === null) {
    window.location.replace('/')
  }
}

const tonConnectUI = new TON_CONNECT_UI.TonConnectUI({
  manifestUrl: 'https://176-99-11-185.cloudvps.regruhosting.ru/tonconnect-manifest.json',
});

async function disconect(){
  await tonConnectUI.disconnect();
  window.location.replace('/')
}

let promo  = document.getElementById('promo')

// add some event listeners to the Element
promo.addEventListener('click', e=> console.log('click!') );
promo.addEventListener('click', e=> console.log('click 2!') );
promo.addEventListener('mouseover', e=> console.log('mouse over!') );

// retrieve the listeners
console.log( promo.getEventListeners() );