from django.shortcuts import render
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from .models import User  # Импорт модели User

async def home(request):
    context = {}
    return render(request, 'home.html', context)

def index(request):
    context = {}
    return render(request, 'index.html', context)


# def home(request):
# async def home(request):
#     context = {}
#     return render(request, 'home.html', context)


# @csrf_exempt
# def update_data(request):
#     if request.method == "POST":
#         key = request.POST.get('key')
#         User.objects.create(key=key)  # Создание новой записи с ключом
#         return JsonResponse({"message": "Key saved successfully"})
#     else:
#         return JsonResponse({"error": "Invalid request"}, status=400)