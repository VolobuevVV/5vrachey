<template>
  <div class="backend-test">
    <h2>Тест подключения к бэкенду</h2>
    <button @click="fetchData">Получить данные с бэкенда</button>
    <div v-if="loading">Загрузка...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else-if="data" class="success">
      <h3>Ответ от бэкенда:</h3>
      <pre>{{ data }}</pre>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BackendTest',
  data() {
    return {
      data: null,
      loading: false,
      error: null
    }
  },
  methods: {
    async fetchData() {
      this.loading = true
      this.error = null
      try {
        // Делаем запрос к нашему Go-бэкенду
        const response = await fetch('http://localhost:8080/api')
        if (!response.ok) {
          throw new Error(`Ошибка HTTP: ${response.status}`)
        }
        this.data = await response.json()
      } catch (err) {
        this.error = err.message
        console.error('Ошибка при запросе к бэкенду:', err)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.backend-test {
  margin: 20px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

button {
  padding: 10px 20px;
  background: #42b883;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background: #369870;
}

.error {
  color: #ff4757;
  margin-top: 10px;
}

.success {
  color: #2ed573;
  margin-top: 10px;
}

pre {
  background: #f4f4f4;
  padding: 10px;
  border-radius: 4px;
  overflow-x: auto;
}
</style>
