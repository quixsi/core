<template>
  <div class="max-w-2xl mx-auto mt-12 flex flex-col gap-2">
    <h3 class="text-lg font-semibold mb-2">Create User</h3>
    <div v-if="error" class="text-red-500 text-sm mt-1">{{ error }}</div>
    <div class="flex gap-2">
        <Input
            v-model="newName"
            placeholder="Name"
            class="flex-1"
        />
        <Input
            v-model="newEmail"
            placeholder="Email"
            class="flex-1"
        />
    </div>
    <Button 
        class="bg-slate-500 hover:bg-blue-400 text-white"
        @click="createNewUser"
    >
        Create new user
    </Button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Input, Button } from '@quixsi/components'
import { useUsers } from '../useUsers'

const newName = ref('')
const newEmail = ref('')
const { createUser, error } = useUsers()

const createNewUser = async () => {
    await createUser(newName.value, newEmail.value)
}
</script>
