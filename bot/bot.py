import logging
import asyncio
import sys

from aiogram import Bot, Dispatcher
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart
from aiogram.types import Message, WebAppInfo
from aiogram.utils.keyboard import InlineKeyboardBuilder



TOKEN = '7120451302:AAFmiJriqQlj0aDohxsBjIRdTQbbJOMEyhs'

logger = logging.getLogger(__file__)
dp = Dispatcher()
bot = Bot(TOKEN, parse_mode=ParseMode.HTML)

@dp.message(CommandStart())
async def command_start_handler(message: Message):
    mk_b = InlineKeyboardBuilder()
    mk_b.button(text='Site', web_app=WebAppInfo(url='https://176-99-11-185.cloudvps.regruhosting.ru/'))
    await message.answer(text='Click to button below:', reply_markup=mk_b.as_markup())

async def main() -> None:
    await bot.delete_webhook(drop_pending_updates=True)
    await dp.start_polling(bot)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    asyncio.run(main())
