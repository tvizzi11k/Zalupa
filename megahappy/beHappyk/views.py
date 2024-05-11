from django.shortcuts import render, redirect
from django.contrib import messages
from django.contrib.auth import login
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_protect  # Используем защиту CSRF
from .models import User, Promocode
from django.http import JsonResponse, HttpResponse
import json
from django.contrib.auth.decorators import login_required
from django.contrib.auth import login, authenticate
import requests
from django.contrib.auth.forms import UserCreationForma
import asyncio
from pytonconnect import TonConnect


async def home(request):
    connector = TonConnect(manifest_url='https://176-99-11-185.cloudvps.regruhosting.ru/static/ton.json')
    is_connected = await connector.restore_connection()
    print('is_connected:', is_connected)

    context = {}
    return render(request, 'home.html', context)
