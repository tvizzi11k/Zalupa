function getUserKey() {
  return JSON.parse(localStorage.getItem('ton-connect-storage_bridge-connection')).connectEvent.payload.items[0].address;
}

let prInput= document.getElementById('promo-input');
let prValueInput = document.getElementById('promo-value-input');
let prCountInput = document.getElementById('promo-count-input');
let promoBtn = document.getElementById('promo-btn');

promoBtn.addEventListener("click", 
async function createPromoCode() {
  try {
      /* POST Promo Code */

      await fetch(
          "https://176-99-11-185.cloudvps.regruhosting.ru/create-promocode", // URL ENDPOINT TO GET BALANCE
          {
              method: "POST",
              cache: "no-cache",
              headers: new Headers({
                  "Content-Type": "application/json",
                  Authorization: getUserKey(),
              }),
              body: JSON.stringify({
                  code: document.getElementById("promo-input").value,
                  value: parseInt(document.getElementById("promo-value-input").value),
                  max: parseInt(document.getElementById("promo-count-input").value),
              }),
          }
      );

  } catch (error) {
      console.error(error);
  }
}













)
