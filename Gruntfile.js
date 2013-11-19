module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    exec: {
      test: {
        cmd: 'go test'
      },

      setup_test: {
        cmd: 'node test_setup.js'
      },
    },

    watch: {
      scripts: {
        files: ['**/*.go'],
        tasks: ['exec:setup_test', 'exec:test'],
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
