<script setup lang="ts">
import { RouterView } from 'vue-router'
import { Toast } from '@quixsi/components'
import { useToast } from '@/composables/useToast'

const { toast } = useToast()
</script>

<template>
  <Toast 
      v-if="toast.message" 
      :message="toast.message" 
      :type="toast.type" 
      :toastKey="toast.key" 
      @click="toast.message = ''"
  />
  <RouterView />
</template>
