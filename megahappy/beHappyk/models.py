from django.db import models
from django.contrib.auth.models import AbstractBaseUser

# Create your models here.
class User(models.Model):
    telegram_id = models.CharField(max_length=255, unique=True, null=True)
    username = models.CharField(max_length=255)
    balance = models.DecimalField(max_digits=10, decimal_places=2, default=0)
    
    USERNAME_FIELD = 'username'

class Promocode(models.Model):
    code = models.CharField(max_length=255, unique=True)
    value = models.DecimalField(max_digits=10, decimal_places=2)
    used = models.BooleanField(default=False)
