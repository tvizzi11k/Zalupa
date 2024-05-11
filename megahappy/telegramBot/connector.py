from pytonconnect import TonConnect
from tc_storage import TcStorage

MANIFEST = 'https://176-99-11-185.cloudvps.regruhosting.ru/static/ton.json'

def get_connector(chat_id: int):
    return TonConnect(MANIFEST, storage=TcStorage(chat_id))
