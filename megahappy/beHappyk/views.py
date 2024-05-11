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


# def registration(request):
#     if request.user.is_authenticated:
#         return redirect('home')
#     if request.method == 'GET':
#         return render(request, 'registration.html')
#     else:
#         return HttpResponse('Method Not Allowed', status=405)

@csrf_protect
def home(request):
    # if not request.user.is_authenticated:
    #     return redirect('registration')
    if request.method == 'POST':
        promocode_input = request.POST.get('promocode')
        try:
            promocode = Promocode.objects.get(code=promocode_input, used=False)
            user = request.user
            user.balance += promocode.value
            user.save()
            promocode.used = True
            promocode.save()
            messages.success(request, f"Промокод на сумму {promocode.value} применен")
        except Promocode.DoesNotExist:
            messages.error(request, "Промокода не существует или уже был использован.")
    return render(request, 'home.html')

# @csrf_protect
# def telegram_callback(request):
#     code = request.GET.get('code')
#     if code:
#         response = requests.post('https://oauth.ton.org/token', data={
#             'code': code,
#             'client_id': 'YOUR_CLIENT_ID',
#             'client_secret': 'YOUR_CLIENT_SECRET',
#             'redirect_uri': 'https://yourwebsite.com/telegram/callback/',
#             'grant_type': 'authorization_code'
#         }) 
#         if response.status_code == 200:
#             data = response.json()
#             access_token = data.get('access_token')
#             if access_token:
#                 user_info_response = requests.get('https://api.ton.org/getUserInfo', headers={'Authorization': 'Bearer ' + access_token})
#                 if user_info_response.status_code == 200:
#                     user_info = user_info_response.json()
#                     telegram_id = user_info.get('id')
#                     username = user_info.get('username', f'user_{telegram_id}')  

#                     user, created = User.objects.update_or_create(
#                         telegram_id=telegram_id,
#                         defaults={'username': username, 'first_name': user_info.get('first_name', ''), 'last_name': user_info.get('last_name', '')}
#                     )
                    
#                     login(request, user)
#                     return redirect('home')
#                 return JsonResponse({'status': 'error', 'message': 'Failed to get user information'})
#             return JsonResponse({'status': 'error', 'message': 'Failed to get access token'})
#         return JsonResponse({'status': 'error', 'message': 'Error in authentication process', 'details': response.json()})
#     return JsonResponse({'status': 'error', 'message': 'No code provided'})




