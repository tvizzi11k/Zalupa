#!/usr/bin/env python
"""Django's command-line utility for administrative tasks."""
import os
import sys
import asyncio
from pytonconnect import TonConnect


async def a():
    connector = TonConnect(manifest_url='https://176-99-11-185.cloudvps.regruhosting.ru/static/ton.json')
    is_connected = await connector.restore_connection()
    print('is_connected:', is_connected)

def main():
    """Run administrative tasks."""
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'happy.settings')
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line(sys.argv)


if __name__ == '__main__':
    asyncio.get_event_loop().run_until_complete(a())
    main()
