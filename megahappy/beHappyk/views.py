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
from django.contrib.auth.forms import UserCreationForm

def home(request):
    context = {}
    return render(request, 'home.html', context)
