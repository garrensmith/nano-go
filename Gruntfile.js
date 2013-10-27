module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    exec: {
      test: {
        cmd: 'go test'
      },
    },

    watch: {
      scripts: {
        files: ['**/*.go'],
        tasks: ['exec:test'],
        options: {
          spawn: false,
        },
      },
    },
  });

  // Load the plugin that provides the "uglify" task.
  grunt.loadNpmTasks('grunt-exec');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['watch']);

};
